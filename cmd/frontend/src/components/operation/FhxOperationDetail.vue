<template>
    <div id="top" class="opDetail">        
        <div class="header">
            <h2 class="header-title">{{ opRef.Unit.name }}</h2>
            <p class="header-desc">{{ opRef.Unit.desc }}</p>
        </div>
        <article class="item">
            <div class="detail-item" v-for="pa in opRef.ParamList" :key="pa.id">
                <table>
                    <thead>
                        <tr >
                            <td colspan="4">
                                <h3 class="detail-title">{{ pa.name }} 
                                    <a href="#top" title="Nach oben"><span class="icomoon-arrow-up2"></span></a>
                                </h3>
                                <hr/>
                            </td>
                        </tr>
                        <tr>
                            <td colspan="4"><p class="detail-desc">{{ pa.desc }}</p></td>
                        </tr>                        
                        
                    </thead>
                    <tbody>
                        <FHXValues v-if="pa.value.CV != ''" :values="pa.value"></FHXValues>
                        <FHXStringValues v-if="pa.value.value_string !== ''" :values="pa.value"></FHXStringValues>
                    </tbody>
                </table>
            </div>
        </article>    
    </div>
</template>
<script setup>
    import { useRoute } from 'vue-router'
    import {watch, ref} from 'vue'
    import { loadParamsFromOPId } from '@/models/operations.js'

    import FHXValues from './FHXValues.vue';
    import FHXStringValues from './FHXStringValues.vue';

    let op = {
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
  
    async function loadParamsformId() {
        let id = +route.params.id
        // if (!id) return
        let data = await loadParamsFromOPId(id)
        if (data['error']) {
            console.error("Datenbank fehler: ", data['message'])
            return
        }
        if (data.ok) {
            opRef.value = JSON.parse(data.content)
        } else {
            console.error("Fehler in der DB: ", data.message)
        }
    }
    loadParamsformId()
 

</script>
<style lang='css' scoped>
    .opDetail {
        & .header {
            width: 40%;
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
        }
    }
    table {
        width: 40%;        
        box-shadow: 2px 2px 2px rgba(0,0,0,.2), -2px -2px 2px rgba(0,0,0,.2);
        margin-bottom: .5rem;
        border: none;
        & thead {
            text-align: center;
        }        
    }
    .detail-title {
        padding: 0.5rem 1.5rem;
        position: relative;
        & a{            
            position: absolute;
            right: 1.5rem;
        }     
    }
    .detail-desc {
        text-align: left;
        padding: 0.5rem 1.0rem;
        font-family: monospace;
    }
           
</style>