// IconsRefresh.cpp : définit le point d'entrée pour l'application console.
//

#include "stdafx.h"
#include "shlobj.h"

int wmain(int argc, _TCHAR* argv[])
{
	SHChangeNotify(SHCNE_ASSOCCHANGED, SHCNF_IDLIST, NULL, NULL);
	return 0;
}
