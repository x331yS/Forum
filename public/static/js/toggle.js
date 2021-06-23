function toggle() {
    const sec = document.getElementById('sec');
    const nav = document.getElementById('navigation');
    sec.classList.toggle('active');
    nav.classList.toggle('active');
}
const signUpButton = document.getElementById('signUp');
const signInButton = document.getElementById('signIn');
const container = document.getElementById('container');
signUpButton.addEventListener('click', () => {
    container.classList.add("right-panel-active");
});
signInButton.addEventListener('click', () => {
    container.classList.remove("right-panel-active");
});