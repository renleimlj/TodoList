$(document).ready(function(){
    var num = 1;
    $(".sub-btn").click(function() {
        let data = {  'issue': $('#issues').val() };
        if (data.issue != "") {
            $.post('/add', data, function(res) {
                if (res != undefined) {
                    $('.table tbody').append(
                        `<tr>
                            <th scope="row">${num}</th>
                            <td>${res.issues}</td>
                         </tr>`)
                    num++;
                }
            })
        } else {
            alert("不能添加空的代办事项！");
        }
    })

    $("tbody").click(function(event) {
        let msg = $(event.target).text();
        if (msg != "已完成！") {
            $(event.target).text("已完成！");
            alert("已完成:  " + msg);
        } else {
            alert("该事项已完成");
        }
    })
});