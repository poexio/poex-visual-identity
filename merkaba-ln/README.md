## merkaba-ln

![merkaba](https://raw.githubusercontent.com/poexio/poex-visual-identity/master/merkaba-ln/out.gif?token=AADtJ7lPaJ8XlRwrQUKfaF4VjZgpme3Zks5ZybWWwA%3D%3D)

## Requirements

- Go
- https://github.com/sangaline/svganimator

### Instructions

```sh
go get github.com/fogleman/ln/ln
go build merkaba.go
./merkaba
```

Animated GIF:

```sh
convert -loop 0 -delay 1.5 out0*.png out.gif
```

Animated SVG:

```sh
sed -i 's|<polyline|<polyline style="stroke-width:10;stroke-miterlimit:10;stroke-dasharray:none;stroke-linecap:round"|' out0*.svg
sed -i 's|black|#66615b|' out0*.svg
python svganimator.py -s 0.03 -p 3 out.svg out0*.svg
```
