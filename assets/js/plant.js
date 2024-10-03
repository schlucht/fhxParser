class Plant {
    constructor() {
        this.plantId = "";
        this.plantName = "keine";
        this.load();
    }

    save() {
        localStorage.setItem(this.storeId, JSON.stringify(this));
    }

    load() {
        const store = JSON.parse(localStorage.getItem(this.storeId));
        if (store) {
            this.plantId = store.plantId;
            this.plantName = store.plantName;
        }     
    }

    print() {
        return this.plantId + ' ' + this.plantName;
    }

    hasPlant() {
        return this.plantId !== "";   
    }

    // laden der Anlage aus dem Localstorage
    // speichern der Anlage im Localstorage
    // ändern der Anlage im Localstorage
    // löschen der Anlage im Localstorage
}

export { Plant };