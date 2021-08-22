$folder = Get-Property -Name CustomActionData
cd $folder
python -m pip install grpcio
python -m pip install grpcio-tools
python -m pip install pillow
#python -m pip install tesserocr-2.4.0-cp37-cp37m-win_amd64.whl
python -m pip install tesserocr-2.4.0-cp36-cp36m-win_amd64.whl
