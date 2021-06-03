function ChangeIt() {
    const images = [
        'url("/img/banner1.jpg")',
        'url("/img/banner2.jpg")',
        'url("/img/banner3.jpg")',
        'url("/img/banner4.jpg")',
        'url("/img/banner6.jpg")',
        'url("/img/banner7.jpg")',
        'url("/img/banner8.jpg")'
    ]
    const section = document.querySelector(".image")
    section.style.backgroundImage = "linear-gradient(\n" +
        "            to bottom,\n" +
        "            rgba(255, 255, 255, 0) 70%,\n" +
        "            #333333 100%\n" +
        "    )," + images[Math.floor(Math.random() * images.length)];
}