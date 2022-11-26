# Bridge

## Assumption
Imagine  you are developing a desktop app that can edit pictures
loadPicture(path)
editPicture()
savePicture(path)


You want to let your app can both run on windows and linux

however, api for fetching and saving files does not the same between windows and linux
* windows: 
    open(path) []byte
    save(path, []byte)
* linux: 
    read(path) []byte
    write(path, []byte)