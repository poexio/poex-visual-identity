#!/bin/bash
echo "go get fogleman/ln"
go get github.com/fogleman/ln/ln

echo "build moose.go"
go build -o bin/moose moose.go

echo "generating frames"
mkdir -p out
./bin/moose


# NOT SURE ABOUT NEXT, LET US END THIS

# exit





echo "updating svgs"
cd out
mkdir -p large small favicon
for i in out*.svg; do
    sed 's|<polyline|<polyline style="stroke-width:4;stroke-miterlimit:4;stroke-dasharray:none;stroke-linecap:round"|' $i > large/$i
    sed 's|black|#66615b|; s|<polyline|<polyline style="stroke-width:10;stroke-miterlimit:10;stroke-dasharray:none;stroke-linecap:round"|' $i > small/$i
done
sed 's|<polyline|<polyline style="stroke-width:16;stroke-miterlimit:16;stroke-dasharray:none;stroke-linecap:round"|' out000.svg > favicon/out000.svg
cd ..

if [ ! -d "bin/svganimator" ]; then
    echo "clone sangaline/svganimator"
    git clone --depth 1 git@github.com:sangaline/svganimator.git bin/svganimator
fi

echo "making logo"
cp out/large/out000.svg poex-logo-large.svg
cp out/small/out000.svg poex-logo.svg
convert -density 106 -alpha remove -alpha off out/large/out000.svg -gravity center -extent 512x512 png24:poex-logo-large.png
convert -density 106 out/large/out000.svg -background none -gravity center -extent 512x512 poex-logo-alpha-large.png
convert -density 104 -alpha remove -alpha off out/small/out000.svg -gravity center -extent 512x512 png24:poex-logo.png
convert -density 104 out/small/out000.svg -background none -gravity center -extent 512x512 poex-logo-alpha.png
convert -density 48 -alpha remove -alpha off out/small/out000.svg -gravity center -extent 244x244 png24:apple-touch-icon.png

echo "animating svg logo"
python bin/svganimator/svganimator.py -s 0.03 -p 3 poex-anim-large.svg out/large/out*.svg
python bin/svganimator/svganimator.py -s 0.03 -p 3 poex-anim.svg out/small/out*.svg

echo "animating gif logo"
convert -loop 0 -delay 1.5 out/out0*.png poex-anim.gif
