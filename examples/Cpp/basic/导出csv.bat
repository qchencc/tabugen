echo off

set currentDir=%cd%
cd ..\..\..\
set rootDir=%cd%
cd %currentDir%

set PYTHONPATH=%rootDir%

set taxi_alias=python %rootDir%\taksi\cli.py
set filepath="%currentDir%\..\..\datasheet\����.xlsx"

%taxi_alias%  --parser=excel --parse_files=%filepath% --enable_column_skip --cpp_out=%currentDir%\src\AutogenConfig --with_csv_codegen  --source_file_encoding=gbk --out_data_format=csv --out_data_path=%currentDir%\res 
pause