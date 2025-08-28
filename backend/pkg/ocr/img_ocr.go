package ocr

import "fmt"

var InitPrompt = fmt.Sprintf(`
Analyze this image that contains text: "%s"

This image is from ESCOM Malawi's social media about power outages.

Determine:
1. Is this a load shedding schedule chart?
2. Does it contain a time table?
3. What areas/locations are mentioned?
4. Extract any schedule information

Return JSON:
{
    "is_schedule_chart": true/false,
    "contains_time_table": true/false,
    "detected_areas": ["Area 25", "Kawale", ...],
    "extracted_schedule": {
        "areas": [
            {
                "name": "Area 25",
                "times": ["06:00-10:00", "14:00-18:00"]
            }
        ]
    }
}
`, "ocrText")
