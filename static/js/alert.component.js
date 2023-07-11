const template = document.createElement('template');

template.innerHTML = `
    <style>
        .alert {
            position: relative;
            text-align: center;
            padding: 1rem;
            width: 100%;
            margin: 1rem;      
        }
        .alert.critical {
            background-color: rgba(255,0,0,0.2);
            border: solid 1px red;
            border-radius: 3px;
            color: red;
        }
        .alert.success {
            background-color: rgba(0,255,0,0.2);
            border: solid 1px green;
            border-radius: 3px;
            color: green;
        }
        .alert span {
            position: absolute;
            font-size: 2rem;
            right: 10px;
            top: 3px;
            cursor: pointer;
        }
        .alert span:hover {
            font-size: 2.2rem;
        }

    </style>

    <div id="alert" class="alert">Hier kommt die Meldung!
        <span>x</span>
    </div>
`

class OtsAlertMessage extends HTMLElement {

    shadow = null;

    constructor() {
        super();

        this.shadow = this.attachShadow({ mode: 'open' })
    }

    connectedCallback() {       
        const alert = this.shadow.getElementById('alert');
        alert.classList.add(this.getAttribute('type'));
        this.shadow.appendChild(template.content.cloneNode(true));
        this.shadow.querySelector('span').addEventListener('click', function() {
            alert.style.display = 'none';
        })
    }
}

export { OtsAlertMessage }