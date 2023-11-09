<template>
    <div class="opDetail">
        <div class="header">
            <h2 class="header-title">{{ opRef.Unit.name }}</h2>
            <p class="header-desc">{{ opRef.Unit.desc }}</p>
        </div>
        <article class="item">
            <div class="detail-item" v-for="pa in opRef.ParamList" :key="pa.id">
                <h3 class="detail-title">{{ pa.name }}</h3>
                <p class="detail-desc">{{ pa.desc }}</p>
                <table border>
                    <tr>
                        <th>Min</th>
                        <th>Max</th>
                        <th>Value</th>
                        <th>Unit</th>
                    </tr>
                    <tr>
                        <td>-1000</td>
                        <td>1000</td>
                        <td>250</td>
                        <td>mbar</td>
                    </tr>
                </table>
            </div>
        </article>    
    </div>
</template>
<script setup>
    import { useRoute } from 'vue-router'
    import {watch, ref } from 'vue'
    import { loadParamsFromOPId } from '@/models/operations.js'

    const op = {
        Unit: {
            name:"",
            desc:"",
            id: 0,
            ParamList: [
                { name: "", desc:"" },
            ]
        }
    }
    const route = useRoute()  
    const opRef = ref(op)
    
    watch(async ()=>{       
            let id = +route.params.id
            let data = await loadParamsFromOPId(id)
            opRef.value = JSON.parse(data.content)
    })  

</script>
<style lang='css' scoped>
    .opDetail {
        & .header {
            & .header-title {
                background-color: var(--light-blue);
                color: var(--white);
                padding: .5rem;
                border-radius: .6rem;
            }
            & .header-desc {
                padding: .8rem .5rem;
                font-weight: 300;
            }
        }
        & .item {
            padding: .5rem;
            display: flex;
            flex-direction: column;
            gap: 1rem;
            & .detail-item {
                border-bottom: 1px solid var(--light-gray);
            }
            & table {
                width: 100%;
            }
            & table td {
                text-align: center;
            }
        }
    }
</style>