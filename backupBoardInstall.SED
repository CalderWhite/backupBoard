[Version]
Class=IEXPRESS
SEDVersion=3
[Options]
PackagePurpose=InstallApp
ShowInstallProgramWindow=0
HideExtractAnimation=0
UseLongFileName=0
InsideCompressed=0
CAB_FixedSize=0
CAB_ResvCodeSigning=0
RebootMode=N
InstallPrompt=%InstallPrompt%
DisplayLicense=%DisplayLicense%
FinishMessage=%FinishMessage%
TargetName=%TargetName%
FriendlyName=%FriendlyName%
AppLaunched=%AppLaunched%
PostInstallCmd=%PostInstallCmd%
AdminQuietInstCmd=%AdminQuietInstCmd%
UserQuietInstCmd=%UserQuietInstCmd%
SourceFiles=SourceFiles
[Strings]
InstallPrompt=
DisplayLicense=
FinishMessage=Installed, and ready to use!
TargetName=C:\Users\calder\backupBoardInstall.EXE
FriendlyName=Backup Board
AppLaunched=<None>
PostInstallCmd=move BACKUP~1.exe C:%HOMEPATH%\Desktop\backupboard.exe
AdminQuietInstCmd=
UserQuietInstCmd=
FILE0="backupboard.exe"
[SourceFiles]
SourceFiles0=C:\Users\calder\Desktop\backupboard\bin\
[SourceFiles0]
%FILE0%=
