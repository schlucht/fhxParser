function readTitleFromFhx(text) {
    if (text.length === 0) return 'Keine Daten vorhanden!';
    return text.slice(0, 500);    
}

export { readTitleFromFhx }