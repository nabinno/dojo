---
title: Analyzing Students' Mental Health
tags: analytics,data-engineering
url: https://projects.datacamp.com/projects/1593
---

```sql
-- Start coding here...
SELECT
    stay,
    COUNT(*) AS count_int,
    ROUND(AVG(todep), 2) AS average_phq,
    ROUND(AVG(tosc), 2) AS average_scs,
    ROUND(AVG(toas), 2) AS average_as
FROM public.students
WHERE inter_dom = 'Inter'
GROUP BY stay
ORDER BY stay desc;
```
