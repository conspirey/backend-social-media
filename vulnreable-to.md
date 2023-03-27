!!!Cross-Site Scripting (XSS) Attacks!!!

Dieser Beispielcode ist anfällig für xss-Angriffe, da wir die Namensdaten ohne Codierung oder Maskierung an den Browser weitergeben. 
Wenn wir einen bösartigen Code wie <script>alert(“Ich Bin n Vnl”)</script> weitergeben, 
würde er den JavaScript-Code ausführen und die Warnung „I’m Malicious“ anzeigen.

!!! An den Creator dieses projekts, schau nach ob du dies im "backend" reparieren kannst, Das Frontend ist im normalzustand!
