<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="{{.Description}}">
    <meta name="keywords" content="{{.Keywords}}">
    <meta name="author" content="{{.Author}}">

    <title>{{.Title}}</title>

    <!-- Bootstrap core CSS -->
    <link href="/static/bootstrap/css/bootstrap.min.css" rel="stylesheet"/>
    <link href="/static/css/main.css" rel="stylesheet"/>

    <script src="/static/js/jquery.min.js"></script>
    <script src="/static/bootstrap/js/bootstrap.min.js"></script>

    <!-- HTML5 shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!--[if lt IE 9]>
    <script src="/static/js/html5shiv.js"></script>
    <script src="/static/js/respond.min.js"></script>
    <![endif]-->
</head>

<body>

<!-- Fixed navbar -->
<div id="main-nav" class="navbar navbar-default navbar-fixed-top">
    <div class="container">
        <div class="navbar-header">
            <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="#">RaysGo</a>
        </div>
        <div class="navbar-collapse collapse">
            <ul class="nav navbar-nav navbar-right">
                <li class="active"><a href="/">Home</a></li>
                <li><a href="/about">About</a></li>
                <li><a href="/contact">Contact</a></li>
            </ul>
        </div>
        <!--/.nav-collapse -->
    </div>
</div>

<div id="main-wrapper" class="container">
    <div id="message" class="container">
        <!-- TODO -->
    </div>
    <div class="clearfix"></div>
    <div id="content">
        {{.LayoutContent}}
    </div>
</div>
<!-- /container -->

<div id="footer" class="container">
    <hr>
    Copyright &copy; Raysmond 2014. All Right Reserved. - {{.PageStartTime|loadtimes}}ms
    <span style="float: right;">Theme based on <a href="http://getbootstrap.com/">Bootstrap</a> and powered by <a
            href="https://github.com/astaxie/beego">beego</a></span>
</div>

<!-- Custom JavaScript  -->
<script src="/static/js/main.js"></script>
</body>
</html>
