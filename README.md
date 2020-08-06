# demo-go-RestFull

1. install go: https://golang.org/dl/

2. install vscode: https://code.visualstudio.com/download

3. install dogo: go get github.com/liudng/dogo

    dogo ทำหน้าที่ Monitoring การเปลี่ยนแปลงของ Source Code และจะทำการ Compile + Run ให้เราอัตโนมัติ

4. สร้างไฟล์ dogo.json โดยมีรายละเอียดตามนี้
{
    "WorkingDir": ".",
    "SourceDir": [
        "."
    ],
    "SourceExt": [".c", ".cpp", ".go", ".h"],
    "BuildCmd": "go build main.go",
    "RunCmd": "./main",
    "Decreasing": 1
}

... อธิบายคำสั่ง ...

    WorkingDir: working directory, dogo will auto change to this directory.

    SourceDir: the list of source directories.

    SourceExt: monitoring file type.

    BuildCmd: the command of build and compile.

    RunCmd: the program (full) path.

    Decreasing: Ignore the number of modifies, it’s only start counting after build success. Now it’s supported in linux and windows.

