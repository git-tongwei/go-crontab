$(function () {
    // 上传图片点击事件
    $(".upImg").click(function () {
        var fileInput = $(this).children();
        return fileInput[0].click();
    })

    // 图片下文件流改变
    $(".inputFile").on("change", function () {
        var inputFile = $(this);
        var file = inputFile[0].files[0];
        var fileName = file.name;

        var extStart = fileName.lastIndexOf(".");
        var ext = fileName.substring(extStart, file.length).toUpperCase();
        // if (ext != ".PNG" && ext != ".GIF" && ext != ".JPG" && ext != ".JPEG") {
        //     alert('请上传正确的图片');
        //     $(this).val("");
        //     return;
        // }
        // if (file.size > 5242880) {
        //     alert('图片不能大于5MB');
        //     return;
        // }
        var fileData = new FormData();

        fileData.append("file", file);
        // 上传图片
        $.ajax({
            url: '/server/upload',
            data: fileData,
            processData: false,
            contentType: false,
            type: "post",
            dataType: "json",
            success: function (data) {
                if (data.status == 1) {
                    inputFile.next().val(data.img);
                } else {
                    alert(data.msg);
                }
            },
            error: function (XMLHttpRequest, textStatus, errorThrown) {
                alert('上传失败！');
            },
        })

    })

})