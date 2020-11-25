<!DOCTYPE html>
<head>
    <meta charset="UTF-8">
    <title>接口测试工具</title>
    <script src="../static/js/jquery.min.js"></script>
    <script src="../static/bootstrap/js/bootstrap.js"></script>
    <link rel="stylesheet" href="../static/bootstrap/css/bootstrap.css" />
</head>
<body>
<div class="container-fluid">
    <div class="row" style="background-color: darkslategray;height: 80px;">
        <div class="col-md-24" style="margin-left: 20px;"><h3 style="color: white">接口测试工具</h3></div>
    </div>
    <div class="row">
        <div class="panel panel-info">
            <div class="panel-heading">批量查询结果</div>
            <div class="panel-body">
               <!-- <div if="${resultFileName} != null">
                    <span style="wid 200px;">接口响应数据：</span>
                    <a href="@{'/downloadFile/'+${resultFileName}}" text="${resultFileName}"/>
                    <br/>
                   <span style="wid 200px;">数据解析文件：</span>-->
                    <!--<a href="@{'/downloadFile/'+${resultXlsFileName}}" text="${resultXlsFileName}"/>-->
                <!--</div>-->
                <div style="color: #a94442;font-style: oblique">
                    <span style="width:200px;">处理结果：</span>
                    <span text="{{.Msg}}">{{.Msg}}</span>
                    <br/>
                </div>
                <div >
                    <span style="width:200px;">接口响应数据：</span>
                    <a href="../static/file/download/{{.CsvFileName}}" text="{{.CsvFileName}}">{{.CsvFileName}}</a>
                </div>
                <div >
                    <span style="width:200px;">数据解析文件：</span>
                    <a href="../static/file/download/{{.XlsFileName}}" text="{{.XlsFileName}}">{{.XlsFileName}}</a>
                </div>

            </div>
        </div>

    </div>
</div>
</body>
<script>

</script>
</html>