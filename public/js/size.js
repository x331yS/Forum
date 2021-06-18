if (!(/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent))) {
    var script1= document.createElement('script');
    script1.src = 'js/moove.js';
    script1.type = 'text/javascript';
    document.getElementsByTagName('body')[0].appendChild(script1);
}
