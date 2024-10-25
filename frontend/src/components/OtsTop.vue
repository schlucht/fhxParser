<template>
    <div class="top-plant">
        <h3 id="plant">{{ plant }}</h3>
    </div>
</template>
<script setup>
    import { store } from '../store/store.js'; 
    import { onMounted, ref } from 'vue';

    const plant = ref("");
    if (store.plant.value) {
        plant.value = store.plant.value.plant;    
    }
   
</script>
<style lang='css' scoped>
.top-plant {
    display: flex;
    justify-content: right;
    align-items: center;
    height: 100%;
    width: 100%;
    font-size: 2rem;
    color: var(--clr-white);
    background-color: var(--clr-blue);
}

.no-plant {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    position: absolute;
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
    background-color: rgba(0,0,0,0.6);
    color: var(--clr-orange);
    & h3 {
        margin-bottom: 2rem;
    }
    & .btn {
        width: 12rem;
    }
}

.no-plant-select {
    color: var(--clr-white);
    font-size: 2rem;
    border: solid var(--clr-white) .2rem;
    padding: 3rem 1rem;
    margin-bottom: 2rem;
    & input {
        margin-left: 2rem;
    }
}
</style>