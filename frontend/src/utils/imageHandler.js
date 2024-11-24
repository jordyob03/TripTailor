export const convertFileToBase64 = (file) => {
    return new Promise((resolve, reject) => {
        const reader = new FileReader();
        reader.onloadend = () => {
            const base64String = reader.result.split(",")[1]; 
            resolve(base64String);
        };
        reader.onerror = reject;
        reader.readAsDataURL(file); 
    });
};

export const convertBase64ToImage = (base64String, mimeType = 'image/jpeg') => {
    const byteCharacters = atob(base64String); 
    const byteNumbers = new Array(byteCharacters.length);
    for (let i = 0; i < byteCharacters.length; i++) {
        byteNumbers[i] = byteCharacters.charCodeAt(i);
    }
    const byteArray = new Uint8Array(byteNumbers);
    const blob = new Blob([byteArray], { type: mimeType }); 
    return URL.createObjectURL(blob); 
};
