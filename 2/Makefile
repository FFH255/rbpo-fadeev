MAIN_PKG = src/cmd/bruteforce/main.go

BUILD_DIR = dist

CONFIG_PATH = config.yaml

serve:
	go build -o ${BUILD_DIR}/main.exe ${MAIN_PKG}
	${BUILD_DIR}/main.exe -config=${CONFIG_PATH}

build: 
	go build -o ${BUILD_DIR}/main.exe ${MAIN_PKG}
