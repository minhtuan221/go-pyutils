from distutils.dir_util import copy_tree
import os

x = dict()
# x.


def ExportGo(Gopath="/Users/minhtuannguyen/go/src/", excludes=[]):
    # Get current dir
    fromDirectory = os.getcwd()
    # get current dir name
    foldername = os.path.basename(fromDirectory)

    # copy subdirectory example
    # combine current dir with go-path
    toDirectory = str(Gopath)+foldername

    # Get all subdir abs path
    from glob import glob
    listdir = glob("./*/")
    # print(listdir)
    for folder in listdir:
        if folder not in excludes:
            copy_tree(fromDirectory + folder[1:-1], toDirectory+folder[1:-1])
            print("export: ", folder)

# d = '.'
# listFolder = [os.path.join(d, o) for o in os.listdir(d) if os.path.isdir(os.path.join(d, o))]
# print(listFolder)
def main():
    ExportGo(excludes=['./test2/'])

if __name__ == '__main__':
    main()