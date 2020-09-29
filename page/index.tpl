<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
{{.}}
    <div>
        <form action="/test" method="post">
        用户名：<input name="user" type="text"/></br>
        <input type="submit" value="提交"/>
        </form>
    </div>
    <div>
        <form action="/upload" method="post" enctype="multipart/form-data">
        文件：<input name="ufile" type="file"/></br>
        <input type="submit" value="提交"/>
        </form>
     </div>
</body>
</html>