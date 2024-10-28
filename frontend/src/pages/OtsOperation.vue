<template>
    <div class="operation w3-container">
        <div class="w3-row">
            <h1 class="w3-col l12 w3-center">Operationen</h1>            
        </div>
        <div class="w3-row">
            <div class="operation__names w3-col l4">
                <h3 class="operation__subtitle  w3-center" data-content="10" >OP Auflistung</h3>
                <div class="operation__search">                
                    <input id="operation-search" class="operation__search__input" type="text" placeholder="Suchbegriff eingeben">
                    <button class="operation__search__label w3-button"><i class="icomoon-search2"></i></button>
                </div>
                <div class="w3-container operation__items">
                    <div class="operation__item" v-for="op of ops" :key="op.op_id">
                        <a href="javascript:void{0}" class="operation__item__link">{{ op.op_name }}</a>
                    </div> 
                </div>  
            </div>
            <div class="operation__details w3-col l8">
                <h3 class="operation__subtitle w3-center">OP_HKK</h3>
                <div class="w3-container">
                    <p> Hier kommt dann die neue Tabelle </p>
                </div>
            </div>
        </div>
    </div>
</template>
<script setup>
import { loadOperation } from '../models/operation.js';
import { store } from '../store/store.js';
import { ref } from 'vue';
import notie from 'notie';

const ops = ref(null);
const plantId = store.plant.id;

async function load() { 
    const data = await loadOperation(store.plant.id);
    if (data.error) {
        console.log(data)
        notie.alert({
            type: 'error',
            text: data.message,
            timeout: 5000,
        });
        return;
    }    
    ops.value = data.data;
}

load()
</script>
<style lang='css' scoped>

.operation__names {
    border-right: solid 1px rgba(0,0,0,0.3);
    padding: 0.5rem
    
}
 .operation__items{
    display: flex;
    flex-wrap: wrap;
    gap: 1rem; 
    margin-top: 1rem;   
 }
 .operation__item {
    background: var(--clr-blue);
    color: white;
    padding: 1rem;
    border-radius: .3rem;
    cursor: pointer;
 }
 .operation__item__link{
    text-decoration: none;
    font-weight: 400;
 }

 .operation__subtitle {
    position: relative;
    margin: 0.5rem 1rem 1rem 1rem;
    background-color: var(--clr-blue);
    color: white;
    box-shadow: 2px 2px 2px rgba(0,0,0,0.3);
    padding: 0.3rem
 }
 .operation__subtitle[data-content]::after {
    position: absolute;
    content: attr(data-content);
    top: 5px;
    bottom: 5px;
    right: 10px;   
    width: 4rem;
    background: var(--clr-orange);
    border-radius: 50%;
 }

.operation__search {
    position: relative;    
    height: 3rem;    
}

.operation__search__input {
    position: absolute;
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
    appearance: none;
    border: none;
    margin: 0 1rem 0rem 1rem;
    padding-left: 0.5rem;
    box-shadow: 1px 1px 1px rgba(0,0,0,0.5), -1px -1px 1px rgba(0,0,0,0.5);
}

.operation__search__label {
    position: absolute;
    top: 0;
    bottom: 0;
    right: 1rem;
    font-size: 1.8rem;
    line-height: 1.8rem;    
    color: black;    
}
</style>
