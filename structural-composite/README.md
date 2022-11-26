# Composite

## Assumption
Imagine  you have a folder, there are a few subfolders and files in it, and some files may be contained into the subfolders, etc

the client want to caculate how many files in a folder, contains the files in its subfolder

Folder:
    Count() int
    Add(Folder/File)

File:
    Count() => 1