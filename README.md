# To Do List

----

## 效果图

![](https://s1.ax1x.com/2017/11/20/2VXh4.png)

## 支持静态文件服务

可以加载assets目录下的文件。如效果图，加载了样式、图片。

## 支持简单 js 访问

可以加载assets目录下js的文件。如效果图，点击添加会增加一行待办事项，点击待办事项，会显示已完成!还添加了一些判断，如添加空事项会弹出信息并添加失败。

## 提交表单，并输出一个表格

提交一个待办事项，会向表格中按序添加一行。

## 对 /unknown 给出开发中的提示，返回码 5xx

当在搜索框中输入`http://localhost:3000/unknown`会显示`/unknown Not Found 500`，并且在终端中显示

    [martini] Started GET /unknown for [::1]:61025
    [martini] Page Not Found[500]
