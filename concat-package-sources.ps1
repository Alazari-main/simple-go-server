$root = Get-ChildItem "./" -Directory -Recurse
foreach($dir in $root) {

    $resultFilePath = $dir.FullName + [System.IO.Path]::DirectorySeparatorChar + "package.source.go"
    $inputFilesPattern = $dir.FullName + ([System.IO.Path]::DirectorySeparatorChar) + "*.go"

    [System.IO.File]::Delete($resultFilePath)

    $fileNames = (Get-ChildItem -Path $dir.FullName -File -Filter "*.go").FullName

    if(!$fileNames) {
        continue
    }

    Get-Content $fileNames | Out-File $resultFilePath

}
