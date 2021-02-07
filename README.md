# 知苗易约抢购小程序

 **因技术原因尚未完成，遇到了一些问题，希望有人能给出建议**

**建了个群，感兴趣的可以加一下：788512807**

* 问题进度
  * **【已解决】获取验证码时，并未得到验证码图片，而是一大串文字**
    * 该大串响应报文是图片的Base64编码，可转换成图片
  * **【已解决】滑块验证码的x轴坐标可通过opencv来解决。附一个大佬的Python代码，亲测可用https://github.com/crazyxw/SlideCrack**
    * 在golang调用Python程序时，若该*.py文件中没有导入第三方依赖，可顺利被go调用。当导入三方依赖时，会报错，这种情况目前本人无法解决。
    * 由于本人是Windows系统，故将python代码打包成了*.exe可执行程序进行调用，测试后可行
  * **【已解决】滑块验证码获取需要一个zftsl的参数，猜测与时间戳和sessionid有关，目前无法得到该值，故无法拿到滑块验证码文本。**
    * 逆向小程序后，拿到相关的js文件，找到了zftsl参数的来源。
    * 在使用golang调用js时，出现了与调用Python一样的情况，即不能有三方依赖。作者通过将几个js文件函数拼接到了一个文件中-app.js。go调用该js后可得到zftsl参数。测试通过。
  * **【已解决】在拿到滑块验证码的Base64编码时，发现该编码太长导致[]byte接收不完全，因而无法复原验证码图片。**
    * 通过将resp.body写入到本地文件中，再用json解析本地文件得到两个图片的base64编码
    * 需要注意的是，在多协程环境下，会有大量的文件读写操作，可能会造成文件的覆盖。未来需要对文件名（如 file+协程ID）进行区分，在验证码识别完成后可不对文件进行删除操作，直接覆盖对应文件名
* 【进行中。。。】目前基本功能已经完成并测试通过，需要调整代码的逻辑结构使得更易用 。
* 【进行中。。。】增加多线程以提高运行效率
* 【进行中。。。】如何设置定时抢购功能


