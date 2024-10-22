class Storage {

    constructor(key, value) {
        this.key = key
        this.value = value
    }
    load() {
        const store = JSON.parse(localStorage.getItem(this.key));
        if (store) {
            this.value = store.value;
        }
    }
    save() {
        localStorage.setItem(this.key, JSON.stringify({value: this.value}));
    }

    edit(value) {
        this.value = value
        this.save()
    }
    toString() {
        return this.value
    }
}

export { Storage }