<template>
    <header>
        <div class="header-bar">
            <div class="header-bar-icons">
                <span class="icomoon-grid" aria-hidden="true"></span>
                <span class="icomoon-user" aria-hidden="true"></span>
                <span class="icomoon-stars" aria-hidden="true"></span>
                <span class="icomoon-pencil7" aria-hidden="true"></span>
                <span class="icomoon-stamp" aria-hidden="true"></span>
                <span class="icomoon-bell2" aria-hidden="true"></span>
            </div>
            <div class="header-bar-search">
                <input class="search-input" type="text" placeholder="Suche starten...">
                <span class="icomoon-search2 search-icon" aria-hidden="true"></span>
            </div>
        </div>
    </header>
</template>
<script setup>
    
</script>
<style scoped>
    .header-bar {
    height: 5rem;
    width: 100%;
    background-color: var(--light-blue);
    color: var(--white);
    display: flex;
    align-items: center;
    justify-content: flex-end;
}

.header-bar-icons span {
    margin: 0.5rem;
}

.header-bar-search {
    position: relative;
    margin: 1.5rem;
}

.header-bar-search .search-input {
    padding: 0.6rem;
    height: 2.8rem;
    width: 25rem;
    border-radius: var(--border-radius);
    border: none;
    outline: none;
    background-color: var(--blue);
    border: 1px rgba(15, 15, 150, .2) solid;
}

.header-bar-search .search-input::placeholder {
    color: var(--white);
}

.header-bar-search span {
    position: absolute;
    right: 0px;
    top: 0rem;
    bottom: 0rem;
    padding: 0.8rem;
    color: var(--white);
    background-color: var(--mid-blue);
    border-left: 1px rgba(200, 200, 200, .3) solid;
    border-radius: .5rem;
}
</style>