

var container = document.getElementsByClassName('deals');

container[0].onclick = function(event) {

 if (!event.target.classList.contains('deals_item-delete')) return;

  //event.target.parentNode.style.display = "none";
  event.target.parentNode.style.minHeight = "0px";
  event.target.parentNode.style.opacity = "0";
  event.target.parentNode.style.scale = "0";
  event.target.parentNode.style.fontSize = "0px";
  event.target.parentNode.style.hight = "0px";
  event.target.parentNode.style.marginBottom = "-" + 53 +"px";
  event.target.parentNode.style.marginTop = "-" + 53 +"px";
//  setTimeout(function() { event.target.parentNode.style.display = "none"}, 1000);
}
