function AjaxSendWithStatusReturn(data){

    $("#LoadingIndicator").show();

    var xhr = new XMLHttpRequest();
    var url = "/api2";
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-type", "application/json");
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var json = JSON.parse(xhr.responseText);
            $("#LoadingIndicator").hide();
            console.log(json.Status);
            return true
            }
      };

        console.log(data);

        xhr.send(data);

  }
