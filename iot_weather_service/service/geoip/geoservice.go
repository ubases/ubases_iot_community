package geoip

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_weather_service/config"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/golang/groupcache/lru"
	"github.com/mholt/archiver/v3"
	geoip2 "github.com/oschwald/geoip2-golang"
)

const (
	DB_URL    = "https://download.maxmind.com/app/geoip_download?license_key=%s&edition_id=GeoLite2-City&suffix=tar.gz"
	CacheSize = 50000
)

var (
	errNotModified = errors.New("unmodified")
)

type GeoServer struct {
	db       *geoip2.Reader
	dbURL    string
	cache    *lru.Cache
	cacheGet chan get
	dbUpdate chan *geoip2.Reader
}

type get struct {
	ip   string
	resp chan *geoip2.City
}

func NewServer(dbFile, license_key string) (server *GeoServer, err error) {
	server = &GeoServer{
		dbURL:    fmt.Sprintf(DB_URL, license_key),
		cache:    lru.New(CacheSize),
		cacheGet: make(chan get, 10000),
		dbUpdate: make(chan *geoip2.Reader),
	}
	var lastModified time.Time
	if dbFile != "" {
		server.db, lastModified, err = readDbFromFile(dbFile)
		if err != nil {
			iotlogger.LogHelper.Errorf("本地ip库文件读取失败,将从web下载ip库.(%s)", err.Error())
			server.db, lastModified, err = readDbFromWeb(server.dbURL, time.Time{})
			if err != nil {
				return nil, err
			}
		}
	} else {
		server.db, lastModified, err = readDbFromWeb(server.dbURL, time.Time{})
		if err != nil {
			return nil, err
		}
	}
	go server.run()
	go server.keepDbCurrent(lastModified)
	return
}

func (server *GeoServer) run() {
	for {
		select {
		case g := <-server.cacheGet:
			if cached, found := server.cache.Get(g.ip); found {
				g.resp <- cached.(*geoip2.City)
			} else {
				city, err := server.lookupDB(g.ip)
				if err != nil {
					iotlogger.LogHelper.Error(err.Error())
				} else {
					server.cache.Add(g.ip, city)
				}
				g.resp <- city
			}
		case db := <-server.dbUpdate:
			if server.db != nil {
				iotlogger.LogHelper.Debug("Closing old database")
				server.db.Close()
			}
			iotlogger.LogHelper.Debug("Applying new database")
			server.db = db
			iotlogger.LogHelper.Debug("Clearing cached lookups")
			server.cache = lru.New(CacheSize)
		}
	}
}

func (server *GeoServer) lookupDB(ip string) (*geoip2.City, error) {
	return server.db.City(net.ParseIP(ip))
}

func (server *GeoServer) keepDbCurrent(lastModified time.Time) {
	for {
		time.Sleep(48 * time.Hour)
		db, modifiedTime, err := readDbFromWeb(server.dbURL, lastModified)
		if err == errNotModified {
			continue
		}
		if err != nil {
			iotlogger.LogHelper.Errorf("Unable to update database from web: %s", err)
			continue
		}
		lastModified = modifiedTime
		server.dbUpdate <- db
	}
}

func readDbFromFile(dbFile string) (*geoip2.Reader, time.Time, error) {
	dbData, err := ioutil.ReadFile(dbFile)
	if err != nil {
		return nil, time.Time{}, fmt.Errorf("Unable to read db file %s: %s", dbFile, err)
	}
	fileInfo, err := os.Stat(dbFile)
	if err != nil {
		return nil, time.Time{}, fmt.Errorf("Unable to stat db file %s: %s", dbFile, err)
	}
	lastModified := fileInfo.ModTime()
	db, err := openDb(dbData)
	if err != nil {
		return nil, time.Time{}, err
	} else {
		return db, lastModified, nil
	}
}

func readDbFromWeb(url string, ifModifiedSince time.Time) (*geoip2.Reader, time.Time, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, time.Time{}, err
	}
	req.Header.Add("If-Modified-Since", ifModifiedSince.Format(http.TimeFormat))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, time.Time{}, fmt.Errorf("Unable to get database from %s: %s", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotModified {
		return nil, time.Time{}, errNotModified
	}
	if resp.StatusCode != http.StatusOK {
		return nil, time.Time{}, fmt.Errorf("unexpected HTTP status %v", resp.Status)
	}
	lastModified, err := getLastModified(resp)
	if err != nil {
		return nil, time.Time{}, fmt.Errorf("Unable to parse Last-Modified header %s: %s", lastModified, err)
	}

	unzipper := archiver.NewTarGz()
	err = unzipper.Open(resp.Body, 0)
	if err != nil {
		return nil, time.Time{}, err
	}
	defer unzipper.Close()
	for {
		f, err := unzipper.Read()
		if err != nil {
			return nil, time.Time{}, err
		}
		if f.Name() == "GeoLite2-City.mmdb" {
			dbData, err := ioutil.ReadAll(f)
			if err != nil {
				return nil, time.Time{}, err
			}
			if err = ioutil.WriteFile(config.Global.Geo.DbPath, dbData, os.ModePerm); err != nil {
				iotlogger.LogHelper.Errorf("更新IP数据库失败:%s", err.Error())
			}
			db, err := openDb(dbData)
			if err != nil {
				return nil, time.Time{}, err
			}
			return db, lastModified, nil
		}
	}
	return nil, time.Time{}, err
}

func getLastModified(resp *http.Response) (time.Time, error) {
	lastModified := resp.Header.Get("Last-Modified")
	return http.ParseTime(lastModified)
}

func openDb(dbData []byte) (*geoip2.Reader, error) {
	db, err := geoip2.FromBytes(dbData)
	if err != nil {
		return nil, fmt.Errorf("Unable to open database: %s", err)
	} else {
		return db, nil
	}
}
