document.addEventListener("DOMContentLoaded", ()=> {
    const loadFile = document.getElementById('formFile');
    loadFile.addEventListener('change', function() {
        const reader = new FileReader();
        reader.onload = () => {
            document.getElementById('textout').textContent = reader.result
        }
        reader.readAsText(this.files[0])
    });
});