<template>
    <h1>Operationen</h1>
    <section class="operation">
        <section class="left-list" id="card-section">
            <fhx-op-card></fhx-op-card>
        </section>    
        <section class="right-list">
            <RouterView />
        </section>
    </section>
</template>
<script setup>
    
import FhxOpCard from '@/components/card/FhxOpCard.vue'
    
    
</script>
<style scoped>
    h1 {
        text-align: center;        
        padding: var(--padding-5) 0;
    }
    .operation {
        display: flex;
        flex-direction: row;        
        gap: 2rem;
        .left-list {
            flex: 1;            
        }
        .right-list {
            flex: 2;            
        }
    }
</style>