function ChangeIt() {
    const images = [
        'url("../static/img/banner1.jpg")',
        'url("../static/img/banner2.jpg")',
        'url("../static/img/banner3.jpg")',
        'url("../static/img/banner4.jpg")',
        'url("../static/img/banner6.jpg")',
        'url("../static/img/banner7.jpg")',
        'url("../static/img/banner8.jpg")'
    ]
    const section = document.querySelector(".image")
    section.style.backgroundImage = "linear-gradient(\n" +
        "            to bottom,\n" +
        "            rgba(255, 255, 255, 0) 70%,\n" +
        "            #333333 100%\n" +
        "    )," + images[Math.floor(Math.random() * images.length)];
}