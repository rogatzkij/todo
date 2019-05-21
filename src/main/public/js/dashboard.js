



// удалить дело
function taskDelete(elem){
    /*ajax*/
    var xhttp;

    if (window.XMLHttpRequest){
        xhttp=new XMLHttpRequest();
    }
    else {
        xhttp=new ActiveXObject("Microsoft.XMLHTTP");
    }

    let taskID = elem.parentNode.parentNode.parentNode.id

    xhttp.open("DELETE","/dashboard?id="+taskID, true);
    xhttp.send();
    
}