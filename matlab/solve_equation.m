function phi = solve_equation(w, h, ni, pi, iter)
    count = 0;
    phi = zeros(h, w);
    tmp = phi;
    while count < iter
        count = count + 1;
        % set boundary conditions
        % Neumann conditions
        tmp(1, :) = tmp(2, :);
        tmp(h, :) = tmp(h-1, :);
    	tmp(:, 1) = tmp(:, 2);
        tmp(:, w) = tmp(:, w-1);
        % Dirchlet conditions
        tmp(ni) = 0;
        tmp(pi) = 1;
        % Convolution
        phi(2:h-1, 2:w-1) = (tmp(2:h-1, 1:w-2)+tmp(2:h-1, 3:w)...
            +tmp(1:h-2, 2:w-1)+tmp(3:h, 2:w-1))/4;        
        tmp = phi;
    end
end
