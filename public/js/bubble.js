const contain = document.querySelector('#sec')
for (let i = 1; i <= 50; i++) {
    const blocks = document.createElement('div')
    blocks.classList.add('block')
    blocks.onclick=generate
    contain.appendChild(blocks)
}
function generate() {
    anime({
        targets: '.block',
        translateX: function () {
            return anime.random(-800, 800)
        },
        translateY: function () {
            return anime.random(-400, 400)
        },
        scale: function () {
            return anime.random(1, 5)
        }
    })
}
generate()