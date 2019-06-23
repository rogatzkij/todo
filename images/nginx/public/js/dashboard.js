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

    location.reload(true);
}

// пометить сделанным
function taskDone(elem){
    /*ajax*/
    var xhttp;

    if (window.XMLHttpRequest){
        xhttp=new XMLHttpRequest();
    }
    else {
        xhttp=new ActiveXObject("Microsoft.XMLHTTP");
    }

    let taskID = elem.parentNode.parentNode.parentNode.id

    xhttp.open("PUT","/dashboard?id="+taskID, true);
    xhttp.send();

    location.reload(true);
}

// добавить дело дело
function taskAdd(){
    /*ajax*/
    var xhttp;

    if (window.XMLHttpRequest){
        xhttp=new XMLHttpRequest();
    }
    else {
        xhttp=new ActiveXObject("Microsoft.XMLHTTP");
    }

    let elemTitle = document.getElementsByClassName('modal-input-title');
    let elemDescr = document.getElementsByClassName('modal-input-description');
    let elemDate = document.getElementsByClassName('modal-input-date');
    
    const pleaseMsg = "Пожалуйста"
    let errMsg = pleaseMsg; 
    if(elemTitle[0].value == ""){
        errMsg += ", введите заголовок";
    }
    if(elemDescr[0].value == ""){
        errMsg += ", введите описание";
    }
    if(elemDate[0].value  == ""){
        errMsg += ", установите дату";
    }
    errMsg +="."

    if(errMsg == pleaseMsg+"."){
        let querry = "/dashboard?title="+elemTitle[0].value+"&description="+elemDescr[0].value+"&date="+elemDate[0].value
        document.getElementById('id01').style.display='none' 
        xhttp.open("PUT",querry , true);    
        xhttp.send();
        location.reload(true);
    }else{
        document.getElementById('errMessage').innerText = errMsg
        document.getElementById('errMessageContainer').style.display='block'
    }
}

// разлогинеться
function Logout(){
    document.cookie = 'session_id=; expires=Thu, 01 Jan 1970 00:00:01 GMT;';
    location.reload(true);
}

function openWindow(){
    
    document.getElementById('id01').style.display='block';   
    document.getElementById('dashboard-content-id').style.webkitFilter = "blur(10px)";

}

function closeWindow(){
    document.getElementById('id01').style.display='none';   
    document.getElementById('dashboard-content-id').style.webkitFilter = "blur(0px)";
}