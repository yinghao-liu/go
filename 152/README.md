# 152

freetype

字体有几种类型

- TrueType（.TTF）

- 光栅字体（.FON）

- 矢量字体（.FON）



字体的大小可以理解为字体的高度（严格意义上并不是），见参考2。



字体的单位有：

- px：pixel，像素，屏幕上显示的最小单位，用于网页设计，直观方便；
- pt：point，磅，是一个标准的长度单位，1pt＝1/72英寸，用于印刷业；

印刷中没有像素的概念，但是计算机系统里，图像都是由像素构成的，物理的尺寸没有绝对的意义，需要根据DPI装换。



DPI（Dots Per Inch）每英寸点数，输出分辨率，针对于输出设备而言的

PPI （Pixels Per Inch）像素密度单位

A 12-point font is **16 pixels tall**. 

12 points = 12/72 logical inch = 1/6 logical inch = 96/6 pixels = 16 pixels This scaling factor is described as 96 dots per inch (DPI).



1磅对应某个DPI下的像素是：（根据单位运算，最后剩下的单位是像素）

(1磅=1/72 inch) * DPI

1英寸 = 2.54 厘米

1磅=0.03527厘米





## reference

1. [.fon文件和.ttf文件的区别](https://www.cnblogs.com/ruiruizhang/archive/2010/01/16/1648988.html)
1. [字号与行高](https://zhuanlan.zhihu.com/p/27381252)
1. [字号pt px inch cm的关系_号数、点数制对照表](http://www.cnprint.org/bbs/thread/73/67855/)