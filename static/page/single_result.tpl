<!DOCTYPE html>
<head>
    <meta charset="UTF-8">
    <title>融数聚接口测试工具</title>
    <script src="../static/js/jquery.min.js"></script>
    <script src="../static/bootstrap/js/bootstrap.js"></script>
    <link rel="stylesheet" href="../static/bootstrap/css/bootstrap.css" />
</head>
<body>
<div class="container-fluid">
    <div class="row" style="background-color: darkslategray;height: 80px;">
        <div class="col-md-24" style="margin-left: 20px;"><h3 style="color: white">融数聚接口测试工具</h3></div>
    </div>
    <div class="row">
        <div class="panel panel-info">
            <div class="panel-heading">企业查询结果</div>
            <div class="panel-body">
                <div style="color: #a94442">
                    <span>响应码：{{.Code}}</span>
                    <span>描述信息：{{.Msg}}</span>
                </div>
                <br/>
                <div>
                    <span>响应数据：{{.Data}}</span>
                </div>
            </div>
        </div>

    </div>
</div>
</body>
<script>
    //var json = document.getElementById('json').innerText
</script>
</html>