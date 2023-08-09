

async function fileDataUpload(data) {
    const result = {};
    const requestOptions = {
        method: 'POST',
        headers: {
        'Accept': 'application/text',
        'Content-Type': 'application/text'
        },
        body: data
    }

    try {
        const res = await fetch(`${import.meta.env.VITE_API_URL}/read-fhx`, requestOptions);
        const data = res.json();
        return data;
    } catch (err) {
        console.error("Error in FileUpload: ", err)
    }
}

export { fileDataUpload }