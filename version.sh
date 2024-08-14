BUILDDATE=$(date +%Y%m%d)
HEADSTR=$(git rev-parse --short HEAD)
IOTVERSION=2.1.0-CE${BUILDDATE}_${HEADSTR}