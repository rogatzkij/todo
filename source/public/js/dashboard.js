



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

    let querry = "/dashboard?title="+elemTitle[0].value+"&description="+elemDescr[0].value+"&date="+elemDate[0].value
   
    document.getElementById('id01').style.display='none' 
    
    xhttp.open("PUT",querry , true);    
    xhttp.send();
    
    location.reload(true);
}