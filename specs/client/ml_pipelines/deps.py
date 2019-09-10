import importlib


def get_package_version(pkg):


    try:
        return pkg.__version__
    except AttributeError:
        pass
    try:
        return str(pkg.version)
    except AttributeError:
        pass
    try:
        return ".".join(map(str, pkg.version_info))
    except AttributeError:
        return None


def get_package_info(pkgs):
    """ get package versions for the passed required & optional packages """

    pversions = []
    for modname in pkgs:


        try:
            mod = importlib.import_module(modname)
            ver = get_package_version(mod)
            pversions.append((modname, ver))
        except Exception:
            pass

    return pversions