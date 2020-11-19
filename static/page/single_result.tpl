<!DOCTYPE html>
<head>
    <meta charset="UTF-8">
    <title>接口测试工具</title>
    <script src="../static/js/jquery.min.js"></script>
    <script src="../static/bootstrap/js/bootstrap.js"></script>
    <link rel="stylesheet" href="../static/bootstrap/css/bootstrap.css" />
    <style>
        pre {outline: 1px solid #ccc; padding: 5px; margin: 5px; }
        .string { color: green; }
        .number { color: darkorange; }
        .boolean { color: blue; }
        .null { color: magenta; }
        .key { color: red; }

    </style>
    <script type="text/javascript">
        function syntaxHighlight(json) {
            if (typeof json != 'string') {
                json = JSON.stringify(json, undefined, 2);
            }
            json = json.replace(/&/g, '&').replace(/</g, '<').replace(/>/g, '>');
            return json.replace(/("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?)/g, function(match) {
                var cls = 'number';
                if (/^"/.test(match)) {
                    if (/:$/.test(match)) {
                        cls = 'key';
                    } else {
                        cls = 'string';
                    }
                } else if (/true|false/.test(match)) {
                    cls = 'boolean';
                } else if (/null/.test(match)) {
                    cls = 'null';
                }
                return '<span class="' + cls + '">' + match + '</span>';
            });
        }
    </script>
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
                <div style="color: #a94442;font-style: oblique">
                    <span>响应码：{{.Code}}</span>
                    <span>描述信息：{{.Msg}}</span>
                </div>
                <br/>
                <!--<div>
                    <span>响应数据：{{.Data}}</span>
                    <pre id="songReqJson"></pre>
                </div>-->
                <!--<pre id="songReqJson"></pre>-->
                <pre id="result"></pre>
            </div>
        </div>

    </div>
</div>

<script type="text/javascript">

    var result = JSON.stringify(JSON.parse({{.Data}}), null, 2);
    document.getElementById('result').innerHTML = syntaxHighlight(result);
</script>
</body>

</html>