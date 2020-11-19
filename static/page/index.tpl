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
        <ul id="myTab" class="nav nav-tabs">
            <li class="active"><a href="#ent" data-toggle="tab">单条查询</a></li>
            <li ><a href="#ents">企业批量查询</a></li>
            <li ><a href="#person">个人批量查询</a></li>
        </ul>
        <div id="myTabContent" class="tab-content">
            <div class="tab-pane fade in active" id="ent">
                <div class="panel panel-info">
                    <div class="panel-heading">单条查询</div>
                    <div class="panel-body">
                        <form class="form-horizontal" id="entform" action="/singleQuery" method="post">
                            <div class="form-group">
                                <label for="mykey" class="col-sm-2 control-label">机构标识</label>
                                <div class="col-sm-4">
                                    <input type="text" class="form-control" id="mykey" name="mykey" placeholder="输入机构标识秘钥串">
                                </div>
                            </div>
                            <div class="form-group">
                                <label for="signkey" class="col-sm-2 control-label">签名秘钥</label>
                                <div class="col-sm-4">
                                    <input type="text" class="form-control" id="signkey" name="signkey" placeholder="输入签名秘钥串">
                                </div>
                            </div>
                            <div class="form-group">
                                <label class="col-sm-2 control-label">测试接口</label>
                                <div class="col-sm-4">
                                    <select class="form-control api-ent" name="apipath">
                                        <option value="rsj/ent/base/get_ent_info">工商照面信息</option>
                                        <option value="rsj/ent/base/get_ent_check">抽查检查</option>
                                        <option value="rsj/ent/base/get_ent_mort">动产抵押类信息</option>
                                        <option value="rsj/ent/case/get_case_punish">工商处罚信息</option>
                                        <option value="rsj/ent/stock/get_ent_stock_impawn">股权出质信息</option>
                                        <option value="rsj/ent/stock/get_ent_sfxz">股权冻结</option>
                                        <option value="rsj/ent/stock/get_ent_sfxz_modify">股权强制转让</option>
                                        <option value="rsj/ent/base/get_ent_certificate">行政许可信息</option>
                                        <option value="rsj/ent/base/get_ent_abnormity">经营异常</option>
                                        <option value="rsj/ent/base/get_ent_annual_report">年报信息</option>
                                        <option value="rsj/ent/base/get_ent_modify">企业变更</option>
                                        <option value="rsj/ent/base/get_ent_filiation">企业分支机构信息</option>
                                        <option value="rsj/ent/base/get_ent_manager">企业高管信息</option>
                                        <option value="rsj/ent/stock/get_ent_inv">企业股东出资信息</option>
                                        <option value="rsj/ent/base/get_ent_liquidation">清算信息</option>
                                        <option value="rsj/ent/case/get_case_detail">信用网站公示处罚信息</option>
                                        <option value="rsj/ent/base/get_ent_yzwfsx">严重违法</option>
                                        <option value="rsj/rsj/ent/base/get_ent_ipr">知识产权出质信息</option>
                                        <option value="rsj/ent/base/get_ent_all">全量查询</option>
                                        <option value="rsj/ent/risk/foreign_lose_info">对外投资企业失信标签</option>
                                        <option value="rsj/ent/risk/equity_lose_info">股东失信标签</option>
                                        <option value="rsj/ent/risk/abnormal_operation_info">经营异常名录标签</option>
                                        <option value="rsj/ent/risk/illegal_tag_info">严重违法失信标签</option>
                                        <option value="rsj/ent/risk/base_tag_info">状态异常标签</option>
                                        <option value="rsj/ent/risk/equity_freeze_info">股权冻结标签</option>
                                        <option value="rsj/ent/risk/all_tag_info">综合风险特征查询标签</option>
                                        <option value="rsj/ent/risk/actual_controller">实际控制人</option>
                                        <option value="rsj/ent/risk/beneficial_owner">受益所有人识别</option>
                                        <option value="rsj/ent/risk/ent_map">企业图谱</option>
                                        <option value="rsj/ent/general/recruitment_info">企业招聘信息</option>
                                        <option value="rsj/ent/general/pop_feel_info">企业舆情信息</option>
                                        <option value="rsj/ent/new_reg_info">新企速递</option>
                                        <option value="rsj/person/card/finance">智慧网点</option>
                                        <option value="rsj/person/passengerstatid/query">个人航空服务报告</option>
                                        <option value="rsj/person/railwaylabel/query">个人铁路服务报告</option>
                                        <option value="rsj/person/network/identity_certification">个人五要素核验</option>
                                        <option value="rsj/person/gsinfo/query">个人工商信息查询</option>
                                        <option value="rsj/person/gsinfo/common">个人工商信息查询(通用)</option>
                                    </select>
                                </div>
                            </div>
                            <div class="form-group">
                                <label for="param1" class="col-sm-2 control-label">参数1</label>
                                <div class="col-sm-4">
                                    <input type="text" class="form-control" id="param1" name="param1" placeholder="输入字段名和值冒号分隔，如 name:张三,可为空">
                                </div>
                            </div>
                            <div class="form-group">
                                <label for="param2" class="col-sm-2 control-label">参数2</label>
                                <div class="col-sm-4">
                                    <input type="text" class="form-control" id="param2" name="param2" placeholder="输入字段名和值冒号分隔，如 name:张三,可为空">
                                </div>
                            </div>
                            <div class="form-group">
                                <label for="param3" class="col-sm-2 control-label">参数3</label>
                                <div class="col-sm-4">
                                    <input type="text" class="form-control" id="param3" name="param3" placeholder="输入字段名和值冒号分隔，如 name:张三,可为空">
                                </div>
                            </div>
                            <div class="form-group">
                                <label for="param4" class="col-sm-2 control-label">参数4</label>
                                <div class="col-sm-4">
                                    <input type="text" class="form-control" id="param4" name="param4" placeholder="输入字段名和值冒号分隔，如 name:张三,可为空">
                                </div>
                            </div>
                            <div class="form-group">
                                <label for="param4" class="col-sm-2 control-label">参数5</label>
                                <div class="col-sm-4">
                                    <input type="text" class="form-control" id="param5" name="param5" placeholder="输入字段名和值冒号分隔，如 name:张三,可为空">
                                </div>
                            </div>
                            <div class="form-group">
                                <div class="col-sm-offset-2 col-sm-4">
                                    <button type="submit" class="btn btn-info">查 询</button>
                                </div>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
            <div class="tab-pane fade" id="ents">
                <div class="panel panel-info">
                    <div class="panel-heading">批量结果查询</div>
                    <div class="panel-body">
                        <form class="form-horizontal" action="/fileQuery" method="post" enctype="multipart/form-data">
                            <div class="form-group">
                                <label class="col-sm-2 control-label">模板下载</label>
                                <div class="col-sm-4">
                                    <a href="../static/file/templates/企业查询模板文件.csv">企业查询模板文件.csv</a>
                                </div>
                            </div>
                            <div class="form-group">
                                <label for="pmykey" class="col-sm-2 control-label">机构标识</label>
                                <div class="col-sm-4">
                                    <input type="text" class="form-control" id="pmykey" name="pmykey" placeholder="输入机构标识秘钥串">
                                </div>
                            </div>
                            <div class="form-group">
                                <label for="psignkey" class="col-sm-2 control-label">签名秘钥</label>
                                <div class="col-sm-4">
                                    <input type="text" class="form-control" id="psignkey" name="psignkey" placeholder="输入签名秘钥串">
                                </div>
                            </div>
                            <div class="form-group">
                                <label class="col-sm-2 control-label">测试接口</label>
                                <div class="col-sm-4">
                                    <select class="form-control api-ents" name="apipath">
                                        <option value="rsj/ent/base/get_ent_info">工商照面信息</option>
                                        <option value="rsj/ent/base/get_ent_check">抽查检查</option>
                                        <option value="rsj/ent/base/get_ent_mort">动产抵押类信息</option>
                                        <option value="rsj/ent/case/get_case_punish">工商处罚信息</option>
                                        <option value="rsj/ent/stock/get_ent_stock_impawn">股权出质信息</option>
                                        <option value="rsj/ent/stock/get_ent_sfxz">股权冻结</option>
                                        <option value="rsj/ent/stock/get_ent_sfxz_modify">股权强制转让</option>
                                        <option value="rsj/ent/base/get_ent_certificate">行政许可信息</option>
                                        <option value="rsj/ent/base/get_ent_abnormity">经营异常</option>
                                        <option value="rsj/ent/base/get_ent_annual_report">年报信息</option>
                                        <option value="rsj/ent/base/get_ent_modify">企业变更</option>
                                        <option value="rsj/ent/base/get_ent_filiation">企业分支机构信息</option>
                                        <option value="rsj/ent/base/get_ent_manager">企业高管信息</option>
                                        <option value="rsj/ent/stock/get_ent_inv">企业股东出资信息</option>
                                        <option value="rsj/ent/base/get_ent_liquidation">清算信息</option>
                                        <option value="rsj/ent/case/get_case_detail">信用网站公示处罚信息</option>
                                        <option value="rsj/ent/base/get_ent_yzwfsx">严重违法</option>
                                        <option value="rsj/ent/base/get_ent_ipr">知识产权出质信息</option>
                                        <option value="rsj/ent/base/get_ent_all">全量查询</option>
                                        <option value="rsj/ent/risk/actual_controller">实际控制人</option>
                                        <option value="rsj/ent/risk/beneficial_owner">受益所有人识别</option>
                                        <option value="rsj/ent/risk/ent_map">企业图谱</option>
                                    </select>
                                </div>
                            </div>
                            <div class="form-group">
                                <label for="entfile" class="col-sm-2 control-label">选择文件</label>
                                <div class="col-sm-4">
                                    <input type="file" id="entfile" name="ufile">
                                    <p class="help-block">下载模板文件编辑后上传</p>
                                </div>
                            </div>
                            <div class="form-group">
                                <div class="col-sm-offset-2 col-sm-4">
                                    <button type="submit" class="btn btn-info">查 询</button>
                                </div>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
            <div class="tab-pane fade" id="person">
                <div class="panel panel-info">
                    <div class="panel-heading">个人接口批量查询</div>
                    <div class="panel-body">
                        <form class="form-horizontal" action="/fileQuery" method="post" enctype="multipart/form-data">
                            <div class="form-group">
                                <label class="col-sm-2 control-label">模板下载</label>
                                <div class="col-sm-4">
                                    <a href="../static/file/templates/个人工商信息查询模板文件.csv">个人工商信息查询模板文件.csv</a><br>
                                    <a href="../static/file/templates/个人铁路出行查询模板文件.csv">个人铁路出行查询模板文件.csv</a><br>
                                    <a href="../static/file/templates/个人航空出行查询模板文件.csv">个人航空出行查询模板文件.csv</a><br>
                                    <a href="../static/file/templates/个人五要素核验查询模板文件.csv">个人五要素核验查询模板文件.csv</a>
                                </div>
                            </div>
                            <div class="form-group">
                                <label for="pmykey" class="col-sm-2 control-label">机构标识</label>
                                <div class="col-sm-4">
                                    <input type="text" class="form-control" id="pmykey" name="pmykey" placeholder="输入机构标识秘钥串">
                                </div>
                            </div>
                            <div class="form-group">
                                <label for="psignkey" class="col-sm-2 control-label">签名秘钥</label>
                                <div class="col-sm-4">
                                    <input type="text" class="form-control" id="psignkey" name="psignkey" placeholder="输入签名秘钥串">
                                </div>
                            </div>
                            <div class="form-group">
                                <label class="col-sm-2 control-label">测试接口</label>
                                <div class="col-sm-4">
                                    <select class="form-control api-ents" name="apipath">
                                        <option value="rsj/person/gsinfo/query">个人工商信息</option>
                                        <option value="rsj/person/gsinfo/common">个人工商信息(通用)</option>
                                        <option value="rsj/person/passengerstatid/query">个人航空出行</option>
                                        <option value="rsj/person/railwaylabel/query">个人铁路出行</option>
                                        <option value="rsj/person/network/identity_certification">个人五要素核验</option>
                                    </select>
                                </div>
                            </div>
                            <div class="form-group">
                                <label for="personfile" class="col-sm-2 control-label">选择文件</label>
                                <div class="col-sm-4">
                                    <input type="file" id="personfile" name="ufile">
                                    <p class="help-block">下载模板文件编辑后上传</p>
                                </div>
                            </div>
                            <div class="form-group">
                                <div class="col-sm-offset-2 col-sm-4">
                                    <button type="submit" class="btn btn-info">查 询</button>
                                </div>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </div>

    </div>
</div>
<div class="row" style="margin-left: 20px;">
    <span>提示：</span><br/>
    <span>1.需要正确填写机构秘钥（16位）、签名秘钥（32位），机构存在可用测试量</span><br/>
    <span>2.接口调用成功且命中，将耗费相应接口的可用数据量</span><br/>
    <span>3.使用批量查询前下载模板文件填写数据，单次最大查询量为50条，请按要求正确填写</span><br/>
    <span>4.批量查询结果文件为两份，csv格式存放接口响应的原数据，xls存放解析后数据</span><br/>
</div>
</div>
</body>
<script>
  $('#myTab a').click(function(e){
    $(this).tab('show')
  })

</script>
</html>