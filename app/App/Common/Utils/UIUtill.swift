//
//  UI.swift
//  app
//
//  Created by sunho on 2018/04/29.
//  Copyright © 2018 sunho. All rights reserved.
//

import Foundation
import UIKit

class UIUtill {
    class var black: UIColor {
        return UIColor.black
    }
    
    class var white: UIColor {
        return UIColor.white
    }
    
    class var gray: UIColor {
        return UIColor(rgba: "#BFBFC3")
    }
    
    class var lightGray1: UIColor {
        return UIColor(rgba: "#D7D7DA")
    }
    
    class var lightGray0: UIColor {
        return UIColor(rgba: "#F0F0F0")
    }
    
    class var blue: UIColor {
        return UIColor(rgba: "#006FFF")
    }
    
    class func roundView(_ view: UIView, _ radius: CGFloat = 10) {
        view.layer.cornerRadius = radius
        view.clipsToBounds = true
    }
}