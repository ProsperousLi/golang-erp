toastr.options = {
    "closeButton": false,
    "debug": false,
    "positionClass": "toast-bottom-right",
    "onclick": null,
    "showDuration": "300",
    "hideDuration": "1000",
    "timeOut": "5000",
    "extendedTimeOut": "1000",
    "showEasing": "swing",
    "hideEasing": "linear",
    "showMethod": "fadeIn",
    "hideMethod": "fadeOut"
}


var _form = $("#form"),
    _pageNum = $("#pageNum"),
    _upload = $(".upload"),
    _file = _form.find("input[type=file]"),
    file = _file[0],
    _fileName = $("#fileName");

$("#showtoast").on("click",function(){
    var data = _form.serializeJson();

    var defaults = {srouceType : 'h5'},
        params = {};
    if(data.mode == "page"){
        defaults.pageNum = data.pageNum;
        defaults.pageSize = data.pageSize || 10;
        defaults.pageNum_s = data.pageNum;
        defaults.pageSize_s = data.pageSize || 10;
        defaults.inner_pageNum = data.pageNum;
        defaults.inner_pageSize = data.pageSize || 10
    }

    try{
        params = eval("["+data.params+"]")[0];
    }catch(e){
        var paramsTexts = data.params.replace(/\n/g,"~!&~").split("~!&~");
        for(var i = 0; i < paramsTexts.length; i++){
            var content1 = paramsTexts[i].split("="),
                content2 = paramsTexts[i].split(":");
            if(content1.length > 1){
                let [key,...val] = content1;
                params[key.trim()] = val.join("=").trim();
            }else if(content2.length > 1){
                let [key,...val] = content2;
                params[key.trim()] = val.join(":").trim();
            }else if(content1[0] === "" && content2[0] === ""){
                continue;
            }else{
                toastr.error("congratulations", "内容解析失败");
                return ;
            }
        }
    }
    params = $.extend({},defaults,params);
    if(data.mode == "upload" && file.files.length > 0){
        var formData = new FormData();
        formData.append(data.uploadName,file.files[0]);
        for(var key in params){
            formData.append(key,params[key]);
        }
        uploadAjax(data,formData);
    }else{
        defaultAjax(data,params);
    }


})

var defaultAjax = function(data,params){
    Dialog.loading();
    $.ajax({
        type : data.type,
        url : data.url,
        data : params,
        dataType: 'json',
    }).done(function(res){
        success(res);
    }).fail(function(error){
        fail(error);
    });
}

var uploadAjax = function(data,params){
    Dialog.loading();
    $.ajax({
        type : data.type,
        url : data.url,
        data : params,
        cache: false,
        processData: false,
        contentType: false,
        dataType: 'json',
    }).done(function(res){
        success(res);
    }).fail(function(error){
        fail(error);
    });
}

var success = function(res){
    Dialog.closeLoading();
    var html = ProcessObject(res, 0, false, false, false);
    $("#Canvas").html("<PRE class='CodeContainer'>"+html+"</PRE>");
}

var fail = function(error){
    Dialog.closeLoading();
    toastr.error("congratulations", error.status);
}


$("#cleartoasts").on("click",function(){
    _form[0].reset();
})

$("input[name=mode]").on("change",function(){
    var val = this.value;
    if(val == "default"){
        _upload.addClass("hide");
        _pageNum.addClass("hide");
    }else if(val =="page"){
        _upload.addClass("hide");
        _pageNum.removeClass("hide");
    }else if(val = "upload"){
        _upload.removeClass("hide");
        _pageNum.addClass("hide");
    }
})

_fileName.on("click",function(){
    _file.trigger("click");
})

_file.on("change",function(){
    _fileName.html(this.files[0].name);
})

String.prototype.format = String.prototype.formats = function(args) {
    var result = this;
    if (arguments.length < 1) {
        return result;
    }

    if(!args){
        result = result.replace(new RegExp("{\\w+}",'g'), "");
        return result;
    }

    var data = arguments;
    if (arguments.length == 1 && typeof (args) == "object") {

        data = args;
    }
    for ( var key in data) {
        var value = data[key];
        if (undefined != value) {
            result = result.replace(new RegExp("{" + key + "}",'g'), value);
        }
    }
    return result;
};
