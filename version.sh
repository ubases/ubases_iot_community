BUILDDATE=$(date +%Y%m%d)
HEADSTR=$(git rev-parse --short HEAD)
IOTVERSION=2.0.0-CE${BUILDDATE}_${HEADSTR}