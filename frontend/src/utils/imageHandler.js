export const convertFileToBytes = (file) => {
    return new Promise((resolve, reject) => {
        const reader = new FileReader();
        reader.onloadend = () => resolve(reader.result); 
        reader.onerror = reject;
        reader.readAsArrayBuffer(file);  
    });
};


export const convertBytesToImage = (byteArray, mimeType = 'image/jpeg') => {
    const blob = new Blob([byteArray], { type: mimeType }); 
    return URL.createObjectURL(blob);
};
