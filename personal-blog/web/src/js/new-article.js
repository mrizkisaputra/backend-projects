$(document).ready(function() {

    $("#form-new-article").on("submit", function(e) {
        e.preventDefault(); // mencegah refresh halaman

        let payload = JSON.stringify({
            title: $("[name='title']").val(),
            content: $("[name='content']").val(),
            category: $("[name='category']").val(),
            publish_date: $("[name='publish_date']").val(),
            tags: $("[name='tags']").val()
        })

        $.ajax({
            url: $(this).attr("action"),
            type: $(this).attr("method"),
            data: payload,
            contentType: "application/json"
        }).done(function(res) {
            alert("Article created successfully")
            window.location.href = "/admin"
        }).fail(function(xhr, status, error) {
            if (xhr.status === 500) {
                alert("Server error: " + xhr.responseText);
            }

            if (xhr.status === 400) {
                alert("Bad request: " + xhr.responseText);
            }

            if (xhr.status === 404) {

            }
        })



    })




})