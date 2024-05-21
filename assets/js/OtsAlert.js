const template = document.createElement('template');

template.innerHTML = /*html*/`
    <style>
        :host {
            display: none;
            position: relative;
            text-align: center;
            padding: 1rem;
            width: 100%;
            margin: 1rem;      
        }
        :host(.critical) {
            background-color: rgba(255,0,0,0.2);
            border: solid 1px red;
            border-radius: 3px;
            color: red;
        }
        :host(.success) {
            background-color: rgba(0,255,0,0.2);
            border: solid 1px green;
            border-radius: 3px;
            color: green;
        }
        :host span {
            position: absolute;
            font-size: 2rem;
            right: 10px;
            top: 3px;
            cursor: pointer;
        }
        :host span:hover {
            font-size: 2.2rem;
        }

    </style>

    <p>        
        <slot>
            Meldung anzeigen!
        </slot>
        <span>x</span>
    </p>
`

class OtsAlertMessage extends HTMLElement {
    static get observedAttribute() { return ['hidden']; }

    constructor() {
        super();
        this._shadow = this.attachShadow({ mode: 'open' });
        this._tmp = document.querySelector('ots-alert-message');
        this._internals = this.attachInternals();
    }

    connectedCallback() {
        this._shadow.appendChild(template.content.cloneNode(true))

        const type = this._tmp.getAttribute('type')

        if (type === 'success') {
            this._tmp.classList.add('success')
        } else if (type === 'critical') {
            this._tmp.classList.add('critical')
        } else {
            this._tmp.classList.add('success')
        }

        const close = this.shadowRoot.querySelector('span')

        close.addEventListener('click', (e) => {
            this._tmp.style.display = 'none';
            this._tmp.removeAttribute('hidden');
        });
    }


    get hidden() { console.log(this.hasAttribute('hidden')); return this.hasAttribute('hidden') }
    set hidden(flag) {
        if (flag) {
            this.setAttribute('hidden', '');
        } else {
            this.removeAttribute('hidden');
        }
    }

    attributeChangedCallback(name, oldValue, newValue) {
        console.log(name);
        if (name === 'hidden') {
            this._tmp.style.display = 'none';
            return;
        }
        this._tmp.style.display = 'block';

    }
}
window.customElements.define('ots-alert-message', OtsAlertMessage);

export { OtsAlertMessage }