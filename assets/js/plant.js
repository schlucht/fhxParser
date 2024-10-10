class Plant {
          constructor() {
              this.storeId = "plant";
              this.plantId = "";
              this.plantName = "keine";
              this.load();
          }

          save(id, name) {
              this.plantId = id;
              this.plantName = name;
              localStorage.setItem(this.storeId, JSON.stringify(this));
              return this;
          }

          load() {
              const store = JSON.parse(localStorage.getItem(this.storeId));
              if (store) {
                  this.plantId = store.plantId;
                  this.plantName = store.plantName;
              }
          }

          toString() {
              return this.plantId + ' ' + this.plantName;
          }

          hasPlant() {
              return this.plantId !== "";
          }
      }
