// Helper function to parse time values
const parseTime = (timeString) => {
    const time = parseFloat(timeString);
    if (timeString.endsWith('ms')) {
        return time;
    } else if (timeString.endsWith('s')) {
        return time * 1000;
    }
    return time; // Assume milliseconds if no unit is provided
};
export default parseTime;