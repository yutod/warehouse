# warehouse

## Install
1. Inside warehouse repository root directory
2. Run following command and put binary file on your $PATH
   If `/usr/local/bin` is not on your `$PATH` please make link on different place
```bash
# (bash, zsh)
ln -snf $(pwd)/warehosue /usr/local/bin/warehouse
```
```fish
# (fish)
ln -snf  (pwd)/warehosue /usr/local/bin/warehouse
```

## Initialize
1. Install all required node modules for gui part
```bash
warehouse init
```

2. Build gui part & Start API server for warehouse
```bash
warehouse run
```

## Start serving gui
```bash
cd $(pwd)/gui
./node_modules/.bin/vue-cli-service serve --port 8081
```
